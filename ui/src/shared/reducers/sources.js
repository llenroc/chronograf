export default function sources(state = [], action) {
  switch (action.type) {
    case 'SOURCES_UPDATED': {
      const updatedSources = action.payload.sources;
      return updatedSources;
    }
  }

  return state;
}
