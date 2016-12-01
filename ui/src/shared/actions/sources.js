export function updateSources(sources) {
  return {
    type: 'SOURCES_UPDATED',
    payload: {
      sources,
    },
  };
}

export function removeSource(source) {
  return {
    type: 'SOURCE_REMOVED',
    payload: {
      source,
    },
  };
}

export function addSource(source) {
  return {
    type: 'SOURCE_ADDED',
    payload: {
      source,
    },
  };
}
