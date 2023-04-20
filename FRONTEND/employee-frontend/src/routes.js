import Vue from "vue";
import Router from "vue-router";
import AuthDataService from "./service/AuthDataService";

Vue.use(Router);

const router = new Router({
    mode: 'history',
    routes: [
        {
            path: "/employees",
            name: "Employees",
            component: () => import("./components/Employees"),
            meta: {
                authRequired: 'true'
            }
        },
        {
            path: "/employee/new",
            name: "NewEmployee",
            component: () => import("./components/NewEmployee"),
            meta: {
                authRequired: 'true'
            }
        },
        {
            path: "/employee/:employeeid",
            name: "Employee",
            component: () => import("./components/Employee"),
            meta: {
                authRequired: 'true'
            }
        },
        {
            path: "/trainings",
            name: "Trainings",
            component: () => import("./components/Trainings"),
            meta: {
                authRequired: 'true'
            }
        },
        {
            path: "/training/new",
            name: "NewTraining",
            component: () => import("./components/NewTraining"),
            meta: {
                authRequired: 'true'
            }
        },
        {
            path: "/training/:trainingid",
            name: "Training",
            component: () => import("./components/Training"),
            meta: {
                authRequired: 'true'
            }
        },  
        {
            path: "/",
            component: () => import("./components/Index"),
            meta: {
                authRequired: 'false'
            }
        },
        {
            path: "/signin",
            component: () => import("./components/Signin"),
            meta: {
                authRequired: 'false'
            }
        },
        {
            path: "/signup",
            component: () => import("./components/Signup"),
            meta: {
                authRequired: 'false'
            }
        },
        {
            path: "/unauthorized",
            component: () => import("./components/Unauthorized"),
            meta: {
                authRequired: 'false'
            }
        },
    ]
});

router.beforeEach((to, from, next) => {
    // console.log('to.meta.authRequired ', to.meta.authRequired)
    if(to.meta.authRequired === 'true') {
        var role = window.localStorage.getItem("role");
        var email = window.localStorage.getItem("email");
        const emailid = to.params.email;
        console.log(role)
        console.log('admin')
        if(
            role == 'admin' ||
            emailid === email
        ){
            return next()
        }else{
            router.push('/unauthorized')
        }
    }else{
        return next()
    }
})

export default router;