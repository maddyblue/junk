import React, { Component } from 'react';
import 'tachyons/css/tachyons.css';
import './App.css';

class App extends Component {
	state = {
		messages: [],
		hover: {},
		symbols: {},
		completion: {},
	};
	componentDidMount() {
		this.openWS();
	}
	openWS = () => {
		const loc = 'ws://localhost:8041/ws';
		console.log('connected to', loc);
		const s = new WebSocket(loc);
		s.onmessage = e => {
			const v = JSON.parse(e.data);
			console.log(v);
			switch (v.Typ) {
				case 'state':
					this.setState(v.Msg);
					break;
				case 'hover':
					this.setState({ hover: v.Msg });
					break;
				case 'symbols':
					const symbols = this.state.symbols;
					symbols[v.Msg.Filename] = v.Msg.Symbols;
					this.setState({ symbols: symbols });
					break;
				case 'completion':
					const completion = this.state.completion;
					completion[v.Msg.Filename] = v.Msg.Completion;
					this.setState({ completion: completion });
					break;
				case 'signature':
					debugger;
					break;
				default:
					const msgs = this.state.messages;
					msgs.unshift(v);
					this.setState({ messages: msgs });
					break;
			}
		};
		s.onclose = () => {
			s.close();
			console.log('WS closed; reconnecting');
			setTimeout(this.openWS, 500);
		};
	};
	command(method, id) {
		var body = new FormData();
		body.set('method', method);
		body.set('id', id);
		fetch('/api/command', { method: 'POST', body: body }).catch(alert);
	}
	open(name, line, char) {
		var body = new FormData();
		body.set('name', name);
		body.set('line', line);
		body.set('char', char);
		fetch('/api/open', { method: 'POST', body: body });
	}
	render() {
		if (!this.state.Wins) {
			return <div>loading...</div>;
		}
		const wins = this.state.Wins.map(v => {
			const symbols = (this.state.symbols[v.Info.Name] || []).map(s => (
				<div key={s.name} className="ma3">
					<span
						className="pa1 ba pointer"
						onClick={() =>
							this.open(
								v.Info.Name,
								s.range.start.line + 1,
								s.range.start.character + 1
							)
						}
					>
						{s.name} {s.detail}
					</span>
				</div>
			));
			const completion = (
				this.state.completion[v.Info.Name] || { items: [] }
			).items.map(c => (
				<div key={c.label} className="ma3">
					<span className="pa1 ba">{c.label}</span> {c.detail}
				</div>
			));
			return (
				<div key={v.Info.Name} className="mv3">
					{v.Info.Name}:
					{v.Methods.map(m => (
						<span
							key={m}
							className="ml1 pa1 ba pointer"
							onClick={() => this.command(m, v.Info.ID)}
						>
							{m}
						</span>
					))}
					{this.state.hover.Filename ? (
						<div
							className="bg-dark-blue ma3 pa1"
							dangerouslySetInnerHTML={{
								__html: this.state.hover.HTML,
							}}
						/>
					) : null}
					{completion.length ? <div>Completion: {completion}</div> : null}
					{symbols.length ? <div>Symbols: {symbols}</div> : null}
				</div>
			);
		});
		return (
			<div>
				{wins}
				{this.state.messages.map((m, n) => (
					<div key={n}>
						<hr />
						{JSON.stringify(m)}
					</div>
				))}
			</div>
		);
	}
}

export default App;
