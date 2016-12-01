import React, {PropTypes} from 'react';
import _ from 'lodash';
import {withRouter, Link} from 'react-router';
import {getKapacitor, deleteSource} from 'shared/apis';
import {
  updateSources as updateSourcesAction,
  removeSource as removeSourceAction,
} from 'src/shared/actions/sources';
import {bindActionCreators} from 'redux';
import {connect} from 'react-redux';

export const ManageSources = React.createClass({
  propTypes: {
    location: PropTypes.shape({
      pathname: PropTypes.string.isRequired,
    }).isRequired,
    source: PropTypes.shape({
      id: PropTypes.string.isRequired,
      links: PropTypes.shape({
        proxy: PropTypes.string.isRequired,
        self: PropTypes.string.isRequired,
      }),
    }),
    addFlashMessage: PropTypes.func,
    updateSourcesAction: PropTypes.func.isRequired,
    removeSourceAction: PropTypes.func.isRequired,
  },
  getInitialState() {
    return {
      sources: [],
    };
  },

  componentDidMount() {
    // getSources().then(({data: {sources}}) => {
      // this.setState({sources}, () => {
    this.state.sources.forEach((source) => {
      getKapacitor(source).then((kapacitor) => {
        this.setState((prevState) => {
          const newSources = prevState.sources.map((newSource) => {
            if (newSource.id !== source.id) {
              return newSource;
            }

            return Object.assign({}, newSource, {kapacitor});
          });

          return Object.assign({}, prevState, {sources: newSources});
        });
      });
    });
  },

  handleDeleteSource(source) {
    const {addFlashMessage} = this.props;

    deleteSource(source).then(() => {
      removeSourceAction(source);
      addFlashMessage({type: 'success', text: 'Source removed from Chronograf'});
    }).catch(() => {
      addFlashMessage({type: 'error', text: 'Could not remove source from Chronograf'});
    });
  },

  render() {
    const {sources} = this.state;
    const {pathname} = this.props.location;
    const numSources = sources.length;
    const sourcesTitle = `${numSources} ${numSources === 1 ? 'Source' : 'Sources'}`;

    return (
      <div id="manage-sources-page">
        <div className="enterprise-header">
          <div className="enterprise-header__container">
            <div className="enterprise-header__left">
              <h1>InfluxDB Sources</h1>
            </div>
          </div>
        </div>
        <div className="container-fluid">
          <div className="row">
            <div className="col-md-12">

              <div className="panel panel-minimal">
                <div className="panel-heading u-flex u-ai-center u-jc-space-between">
                  <h2 className="panel-title">{sourcesTitle}</h2>
                  <Link to={`/sources/${this.props.source.id}/manage-sources/new`} className="btn btn-sm btn-primary">Add New Source</Link>
                </div>
                <div className="panel-body">
                  <div className="table-responsive margin-bottom-zero">
                    <table className="table v-center margin-bottom-zero">
                      <thead>
                        <tr>
                          <th>Name</th>
                          <th>Host</th>
                          <th>Kapacitor</th>
                          <th className="text-right"></th>
                        </tr>
                      </thead>
                      <tbody>
                        {
                          sources.map((source) => {
                            return (
                              <tr key={source.id}>
                                <td>{source.name}{source.default ? <span className="label label-primary">Default</span> : null}</td>
                                <td>{source.url}</td>
                                <td>{_.get(source, ['kapacitor', 'name'], '')}</td>
                                <td className="text-right">
                                  <Link className="btn btn-default btn-xs" to={`${pathname}/${source.id}/edit`}>Edit</Link>
                                  <Link className="btn btn-success btn-xs" to={`/sources/${source.id}/hosts`}>Connect</Link>
                                  <button className="btn btn-danger btn-xs" onClick={() => this.handleDeleteSource(source)}>Delete</button>
                                </td>
                              </tr>
                            );
                          })
                        }
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  },
});

function mapStateToProps(state) {
  return {
    sources: state.sources,
  };
}

function mapDispatchToProps(dispatch) {
  return {
    updateSourcesAction: bindActionCreators(updateSourcesAction, dispatch),
    removeSourceAction: bindActionCreators(removeSourceAction, dispatch),
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(withRouter(ManageSources));
