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

namespace re2net
{
    public class RE2
    {
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
        public static string re2post(string re)
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

        static void Main(string[] args)
        {
            Console.WriteLine(RE2.re2post("te+s*t(t|t?)+"));
        }
    }
}
