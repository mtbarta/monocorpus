export default {
    noteFilter(state, getters, rootState) {
        const filter = state.noteFilter
        if (filter.authors.length == 0) {
            filter.authors = [rootState.login.email]
        }
        if (filter.team == null) {
            filter.team = rootState.login.auth.realm
        }
        filter.title = rootState.route.params.titleFilter

        return filter
    }
}