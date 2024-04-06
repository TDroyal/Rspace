

const ModulePagination = {
    state: {  //存全局数据
        home_current_page:1,//首页的分页信息
        user_postlist_page:1,
    },
    mutations: {     //更新state里面的数据
        updateHomeCurrentPage(state, page) {
            state.home_current_page = page
            // console.log("updateHomeCurrentPage", state.home_current_page)
        },

        updateUserPostlistPage(state, page) {
            state.user_postlist_page = page
        },

        logout_page(state) {
            state.home_current_page = 1
            state.user_postlist_page = 1
            // console.log("logout_page", state.home_current_page)
        },
    },
    actions: {      //对state的各种操作

    },
    modules: {
        
    }
}

export default ModulePagination