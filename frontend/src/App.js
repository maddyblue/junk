import React, { Component } from 'react';
import './App.css';
import SortedTable from './SortedTable';
import { Fetch } from './common';

class App extends Component {
	state = {};
	componentDidMount() {
		this.update();
	}
	update = () => {
		Fetch('Avg', data => {
			this.setState({ data: data });
		});
	};
	render() {
		if (!this.state.data) {
			return 'loading...';
		}
		return (
			<div>
				<SortedTable
					name="averages"
					sort="Average"
					headers={[
						{
							name: 'Name',
						},
						{
							name: 'Average',
							desc: true,
							cell: v => v.toFixed(1),
						},
						{
							name: 'Count',
							desc: true,
						},
						{
							name: 'Styles',
							header: 'Style/Region',
							cell: v => v.sort().join(', '),
							cmp: (a, b) =>
								a
									.sort()
									.join(',')
									.localeCompare(b.sort().join(',')),
						},
					]}
					data={this.state.data}
				/>
			</div>
		);
	}
}

export default App;
