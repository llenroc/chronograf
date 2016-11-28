// import {getSources} from '../apis';

export function updateSources(sources) {
  return {
    type: 'SOURCES_UPDATED',
    payload: {
      sources,
    },
  };
}

// TODO: move updateSources from CheckSources into this
// export function fetchSources(sources) {
//   return (dispatch) => {
//
//   }
// }
