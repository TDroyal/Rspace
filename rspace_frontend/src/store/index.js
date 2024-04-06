import { createStore } from 'vuex'

import ModuleUser from './user'
import ModulePagination from './pagination'

export default createStore({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    user: ModuleUser,
    pagination: ModulePagination,
  }
})
