import React, {PropTypes} from 'react';
import AutoRefresh from 'shared/components/AutoRefresh';
import LineGraph from 'shared/components/LineGraph';
import ReactGridLayout, {WidthProvider} from 'react-grid-layout';
const GridLayout = WidthProvider(ReactGridLayout);
import _ from 'lodash';

const RefreshingLineGraph = AutoRefresh(LineGraph);

export const LayoutRenderer = React.createClass({
  propTypes: {
    timeRange: PropTypes.shape({
      defaultGroupBy: PropTypes.string.isRequired,
      queryValue: PropTypes.string.isRequired,
    }).isRequired,
    cells: PropTypes.arrayOf(
      PropTypes.shape({
        queries: PropTypes.arrayOf(
          PropTypes.shape({
            rp: PropTypes.string,
            text: PropTypes.string.isRequired,
            database: PropTypes.string.isRequired,
            groupbys: PropTypes.arrayOf(PropTypes.string),
            wheres: PropTypes.arrayOf(PropTypes.string),
          }).isRequired
        ).isRequired,
        x: PropTypes.number.isRequired,
        y: PropTypes.number.isRequired,
        w: PropTypes.number.isRequired,
        h: PropTypes.number.isRequired,
        i: PropTypes.string.isRequired,
        name: PropTypes.string.isRequired,
      }).isRequired
    ),
    autoRefreshMs: PropTypes.number.isRequired,
    host: PropTypes.string,
    source: PropTypes.string,
  },

  getInitialState() {
    return ({
      layout: _.without(this.props.cells, ['queries']),
    });
  },

  buildQuery(q) {
    const {timeRange, host} = this.props;
    const {wheres, groupbys} = q;

    let text = q.text;

    text += ` where time > ${timeRange.queryValue}`;

    if (host) {
      text += ` and \"host\" = '${host}'`;
    }

    if (wheres && wheres.length > 0) {
      text += ` and ${wheres.join(' and ')}`;
    }

    if (groupbys) {
      if (groupbys.find((g) => g.includes("time"))) {
        text += ` group by ${groupbys.join(',')}`;
      } else if (groupbys.length > 0) {
        text += ` group by time(${timeRange.defaultGroupBy}),${groupbys.join(',')}`;
      } else {
        text += ` group by time(${timeRange.defaultGroupBy})`;
      }
    } else {
      text += ` group by time(${timeRange.defaultGroupBy})`;
    }

    return text;
  },

  generateGraphs() {
    const {autoRefreshMs, source} = this.props;

    return this.props.cells.map((cell) => {
      const qs = cell.queries.map((q) => {
        return Object.assign({}, q, {
          host: source,
          text: this.buildQuery(q),
        });
      });
      return (
        <div key={cell.i}>
          <h2 className="hosts-graph-heading">{cell.name}</h2>
          <div className="hosts-graph graph-panel__graph-container">
            <RefreshingLineGraph
              queries={qs}
              autoRefresh={autoRefreshMs}
            />
          </div>
        </div>
      );
    });
  },

  render() {
    return (
      <GridLayout useCSSTransforms={false} layout={this.state.layout} isDraggable={false} isResizable={false} cols={12} rowHeight={90} width={1200}>
        {this.generateGraphs()}
      </GridLayout>
    );
  },
});

export default LayoutRenderer;
