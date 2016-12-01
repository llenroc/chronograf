export default function sources(state = [], action) {
  switch (action.type) {
    case 'SOURCES_UPDATED': {
      const updatedSources = action.payload.sources;
      return updatedSources;
    }

    // case 'SOURCE_ADDED': {
    //   const {source} = action.payload;
    //
    // }
  }

  return state;
}
