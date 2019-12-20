import React from 'react';
import PropTypes from 'prop-types';
import {
	MultiDropdownList,
	SingleDropdownRange,
	RangeSlider,
} from '@appbaseio/reactivesearch';

const SearchFilters = ({ currentTopics, setTopics, visible }) => (
	<div className={`flex column filters-container ${!visible ? 'hidden' : ''}`}>
		<div className="child m7">
			<MultiDropdownList
				componentId="username"
				dataField="username"
				placeholder="Select username"
				title="Username"
				filterLabel="Username"
				size={100}
				showCount={true}
				showFilter={true}
				loader="Loading ..."
			/>
		</div>
		<div className="child m7">
			<MultiDropdownList
				componentId="hashtags"
				dataField="hashtags"
				placeholder="Select hashtags"
				title="Hashtags"
				filterLabel="Hashtags"
				showFilter={true}
				size={100}
				showCount={true}
				queryFormat="and"
				onValueChange={setTopics}
				loader="Loading ..."
			/>
		</div>
		<div className="child m7">
			<MultiDropdownList
				componentId="place"
				dataField="place"
				placeholder="Select place"
				title="Places"
				filterLabel="Places"
				showFilter={true}
				size={100}
				showCount={true}
				queryFormat="and"
				onValueChange={setTopics}
				loader="Loading ..."
			/>
		</div>
		<div className="child m7">
			<SingleDropdownRange
				componentId="date"
				dataField="date"
				placeholder="Latest tweets"
				title="Last Active"
				filterLabel="Last Active"
				data={[
					{ start: 'now-1M', end: 'now', label: 'Last 30 days' },
					{ start: 'now-6M', end: 'now', label: 'Last 6 months' },
					{ start: 'now-1y', end: 'now', label: 'Last year' },
				]}
			/>
		</div>
		<div className="child m7">
			<SingleDropdownRange
				componentId="date"
				dataField="date"
				placeholder="Tweet created"
				title="Created"
				filterLabel="Created"
				data={[
					{
						start: '2019-01-01T00:00:00Z',
						end: '2019-12-31T23:59:59Z',
						label: '2019',
					},
					{
						start: '2018-01-01T00:00:00Z',
						end: '2018-12-31T23:59:59Z',
						label: '2018',
					},
					{
						start: '2017-01-01T00:00:00Z',
						end: '2017-12-31T23:59:59Z',
						label: '2017',
					},
					{
						start: '2016-01-01T00:00:00Z',
						end: '2016-12-31T23:59:59Z',
						label: '2016',
					},
					{
						start: '2015-01-01T00:00:00Z',
						end: '2015-12-31T23:59:59Z',
						label: '2015',
					},
					{
						start: '2014-01-01T00:00:00Z',
						end: '2014-12-31T23:59:59Z',
						label: '2014',
					},
					{
						start: '2013-01-01T00:00:00Z',
						end: '2013-12-31T23:59:59Z',
						label: '2013',
					},
					{
						start: '2012-01-01T00:00:00Z',
						end: '2012-12-31T23:59:59Z',
						label: '2012',
					},
					{
						start: '2011-01-01T00:00:00Z',
						end: '2011-12-31T23:59:59Z',
						label: '2011',
					},
					{
						start: '2010-01-01T00:00:00Z',
						end: '2010-12-31T23:59:59Z',
						label: '2010',
					},
					{
						start: '2009-01-01T00:00:00Z',
						end: '2009-12-31T23:59:59Z',
						label: '2009',
					},
					{
						start: '2008-01-01T00:00:00Z',
						end: '2008-12-31T23:59:59Z',
						label: '2008',
					},
					{
						start: '2007-01-01T00:00:00Z',
						end: '2007-12-31T23:59:59Z',
						label: '2007',
					},
				]}
			/>
		</div>
		<div className="child m7">
			<RangeSlider
				componentId="nlikes"
				title="Tweet likes"
				dataField="nlikes"
				range={{ start: 0, end: 300000 }}
				showHistogram={false}
				rangeLabels={{
					start: '0 Likes',
					end: '300K Likes',
				}}
				innerClass={{
					label: 'range-label',
				}}
			/>
		</div>
		<div className="child m7">
			<RangeSlider
				componentId="nreplies"
				title="Tweet replies"
				dataField="nreplies"
				range={{ start: 0, end: 300000 }}
				showHistogram={false}
				rangeLabels={{
					start: '0 Replies',
					end: '300K Replies',
				}}
				innerClass={{
					label: 'range-label',
				}}
			/>
		</div>
		<div className="child m7">
			<RangeSlider
				componentId="nretweets"
				title="Re-Tweets"
				dataField="nretweets"
				range={{ start: 0, end: 300000 }}
				showHistogram={false}
				rangeLabels={{
					start: '0 Re-Tweets',
					end: '300K Re-Tweets',
				}}
				innerClass={{
					label: 'range-label',
				}}
			/>
		</div>
	</div>
);

SearchFilters.propTypes = {
	currentTopics: PropTypes.arrayOf(PropTypes.string),
	setTopics: PropTypes.func,
	visible: PropTypes.bool,
};

export default SearchFilters;
