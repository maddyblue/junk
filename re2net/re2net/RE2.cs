/*
 * Copyright (c) 2012 Matt Jibson <matt.jibson@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

// From: http://swtch.com/~rsc/regexp/nfa.c.txt

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Diagnostics;

namespace re2net
{
    public class RE2
    {
        public string Re { get; set; }
        public string Post { get; set; }
        public State Start { get; set; }
        public Methods Method { get; set; }

        public enum Methods {
            NFA,
            DFA,
        };

        public RE2(string re, Methods method = Methods.DFA)
        {
            Re = re;
            Method = method;

            Post = re2post(re);
            Start = post2nfa(Post);
        }

        private class SList
        {
            public HashSet<State> States { get; set; }

            public SList()
            {
                States = new HashSet<State>();
            }

            public void Addstate(IEnumerable<State> states)
            {
                foreach (var s in states)
                    Addstate(s);
            }

            public void Addstate(State s)
            {
                if (s.Type == StateType.Split)
                {
                    s.Out.ForEach(x => Addstate(x));
                    return;
                }

                States.Add(s);
            }

            public SList Step(char c)
            {
                var nl = new SList();

                foreach (var s in States)
                {
                    if (c == s.C)
                        nl.Addstate(s.Out);
                }

                return nl;
            }

            public bool IsMatch()
            {
                return States.Any(x => x.Type == StateType.Match);
            }

            public override bool Equals(object obj)
            {
                var o = obj as SList;
                return States.SetEquals(o.States);
            }

            public override int GetHashCode()
            {
                /* HashSet equivalency is not obvious. Two identical (i.e., SetEquals() returns true)
                 * HashSets have different values for GetHashCode(), so use this simple workaround for now.
                 */
                return States.Count;
            }
        }

        public bool Match(string s)
        {
            if (Method == Methods.NFA)
            {
                var sl = Startlist();

                foreach (var c in s)
                {
                    sl = sl.Step(c);
                }

                return sl.IsMatch();
            }
            else if (Method == Methods.DFA)
            {
                var ds = Startdstate();

                foreach (var c in s)
                {
                    if (!ds.next.ContainsKey(c))
                    {
                        var sl = ds.l.Step(c);
                        ds.next[c] = DState.Get(sl);
                    }

                    ds = ds.next[c];
                }

                return ds.l.IsMatch();
            }

            throw new System.ArgumentException("Unknown method type");
        }

        private SList Startlist()
        {
            var sl = new SList();
            sl.Addstate(Start);
            return sl;
        }

        private class paren
        {
            public int nalt { get; set; }
            public int natom { get; set; }
        }

        /// <summary>
        /// Convert infix regexp re to postfix notation.
        /// Insert . as explicit concatenation operator.
        /// </summary>
        /// <param name="re">input re</param>
        /// <returns>postfix re</returns>
        private static string re2post(string re)
        {
            int nalt = 0;
            int natom = 0;
            string buf = "";
            var dst = new StringBuilder();
            var paren = new List<paren>();
            var p = new paren();
            paren.Add(p);

            foreach (var r in re)
            {
                switch (r)
                {
                    case '(':
                        if (natom > 1)
                        {
                            --natom;
                            dst.Append('.');
                        }
                        p.nalt = nalt;
                        p.natom = natom;
                        p = new paren();
                        paren.Add(p);
                        nalt = 0;
                        natom = 0;
                        break;
                    case '|':
                        if (natom == 0)
                            throw new Exception("unexpected |");
                        while (--natom > 0)
                            dst.Append('.');
                        nalt++;
                        break;
                    case ')':
                        if (p.Equals(paren[0]))
                            throw new Exception("something bad");
                        if (natom == 0)
                            throw new Exception("unexpected )");
                        while (--natom > 0)
                            dst.Append('.');
                        for (; nalt > 0; nalt--)
                            dst.Append('|');
                        paren.Remove(p);
                        p = paren.Last();
                        nalt = p.nalt;
                        natom = p.natom;
                        natom++;
                        break;
                    case '*':
                    case '+':
                    case '?':
                        if (natom == 0)
                            throw new Exception("unexpected " + r);
                        dst.Append(r);
                        break;
                    default:
                        if (natom > 1)
                        {
                            --natom;
                            dst.Append('.');
                        }
                        dst.Append(r);
                        natom++;
                        break;
                }
            }

            if (p.Equals(paren.Last()) != true)
                throw new Exception("unbalanced parens");
            while (--natom > 0)
                dst.Append('.');
            for (; nalt > 0; nalt--)
                dst.Append('|');

            return dst.ToString();
        }

        public enum StateType
        {
            Character,
            Match,
            Split,
        };

        public class State 
        {
            public StateType Type { get; set; }
            public char? C { get; set; }
            public List<State> Out { get; set; }

            public State(StateType t, char? c = null, State o = null, State o1 = null)
            {
                Type = t;
                C = c;
                Out = new List<State>();

                if (o != null)
                    Out.Add(o);
                if (o1 != null)
                    Out.Add(o1);
            }

            public string ToString()
            {
                if (Type == StateType.Character)
                    return string.Format("{0}: {1}", C, Out.Count);
                return string.Format("{0}: {1}", Type, Out.Count);
            }
        }

        public class Frag
        {
            public State Start { get; set; }
            public List<State> Out { get; set; }

            public Frag(State s, State r)
            {
                Start = s;
                Out = new List<State> { r };
            }

            public Frag(State s, List<State> r)
            {
                Start = s;
                Out = r;
            }

            public string ToString()
            {
                return string.Format("{0} -> {1}", Start.ToString(), Out.Count);
            }

            public void Patch(State s)
            {
                Out.ForEach(x => x.Out.Add(s));
            }

            public void Append(IEnumerable<State> e)
            {
                Out.AddRange(e);
            }
        }

        /// <summary>
        /// Convert postfix regular expression to NFA.
        /// </summary>
        /// <param name="postfix">postfix re</param>
        /// <returns>start State</returns>
        private static State post2nfa(string postfix)
        {
            var stack = new Stack<Frag>();
            Frag e1, e2, e;
            State s;

            foreach (var p in postfix)
            {
                switch (p)
                {
                    default:
                        s = new State(StateType.Character, c: p);
                        stack.Push(new Frag(s, s));
                        break;
                    case '.':
                        e2 = stack.Pop();
                        e1 = stack.Pop();
                        e1.Patch(e2.Start);
                        stack.Push(new Frag(e1.Start, e2.Out));
                        break;
                    case '|':
                        e2 = stack.Pop();
                        e1 = stack.Pop();
                        s = new State(StateType.Split, o: e1.Start, o1: e2.Start);
                        e1.Append(e2.Out);
                        stack.Push(new Frag(s, e1.Out));
                        break;
                    case '?':
                        e = stack.Pop();
                        s = new State(StateType.Split, o: e.Start);
                        e.Append(new List<State> { s });
                        stack.Push(new Frag(s, e.Out));
                        break;
                    case '*':
                        e = stack.Pop();
                        s = new State(StateType.Split, o: e.Start);
                        e.Patch(s);
                        stack.Push(new Frag(s, s));
                        break;
                    case '+':
                        e = stack.Pop();
                        s = new State(StateType.Split, o: e.Start);
                        e.Patch(s);
                        stack.Push(new Frag(e.Start, s));
                        break;
                }
            }

            e = stack.Pop();

            if (stack.Count > 0)
                throw new Exception();

            e.Patch(new State(StateType.Match));
            return e.Start;
        }

        private class DState
        {
            private static Dictionary<SList, DState> states { get; set; }

            static DState()
            {
                states = new Dictionary<SList, DState>();
            }

            public SList l { get; set; }
            public Dictionary<char, DState> next { get; set; }

            public DState()
            {
                next = new Dictionary<char, DState>();
            }

            public static DState Get(SList sl)
            {
                if (!states.ContainsKey(sl))
                {
                    states.Add(sl, new DState
                    {
                        l = sl,
                    });
                }

                return states[sl];
            }
        }

        private DState Startdstate()
        {
            return DState.Get(Startlist());
        }

        static void Main(string[] args)
        {
            for (int i = 1; i < 30; i++)
            {
                var resb = new StringBuilder();
                var ssb = new StringBuilder();
                for (int j = 1; j <= i; j++)
                {
                    resb.Append("a?");
                    ssb.Append("a");
                }
                resb.Append(ssb);

                var dfa = new RE2(resb.ToString(), Methods.DFA);
                var nfa = new RE2(resb.ToString(), Methods.NFA);
                var cre = new Regex(resb.ToString());
                var s = ssb.ToString();

                var swnfa = new Stopwatch();
                swnfa.Start();
                var nfab = nfa.Match(s);
                swnfa.Stop();

                var swdfa = new Stopwatch();
                swdfa.Start();
                var dfab = dfa.Match(s);
                swdfa.Stop();

                var swdfa2 = new Stopwatch();
                swdfa2.Start();
                var dfab2 = dfa.Match(s);
                swdfa2.Stop();

                var swcre = new Stopwatch();
                swcre.Start();
                var creb = cre.IsMatch(s);
                swcre.Stop();

                if (nfab == false || dfab == false || dfab2 == false || creb == false)
                    throw new Exception("false RE");

                Console.WriteLine("{0:00}: {1}, {2}, {3}, {4}", i, swnfa.Elapsed, swdfa.Elapsed, swdfa2.Elapsed, swcre.Elapsed);
            }
        }
    }
}
