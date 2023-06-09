import Vue from "vue";
import Router from "vue-router";
import AuthDataService from "./service/AuthDataService";
import EmployeeDataService from "./service/EmployeeDataService";

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
                authRequired: 'false'
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
        {
            path: "/account/:email",
            component: () => import("./components/Account"),
            meta: {
                authRequired: 'false'
            }
        },
        {
            path: "/account/employee/:employeeid",
            redirect: "/employee/:employeeid"
        },
        {
            path: "/edit/avatar/:employeeid",
            component: () => import("./components/Avatar"),
            meta: {
                authRequired: 'true'
            }
        },
        {
            path: "/dashboard",
            component: () => import("./components/TrainingDashboard"),
            meta: {
                authRequired: 'true'
            }
        }
    ]
});

router.beforeEach((to, from, next) => {
    // console.log('to.meta.authRequired ', to.meta.authRequired)
    if(to.meta.authRequired === 'true') {
        var role = JSON.parse(window.localStorage.getItem("user")).role;
        const routeEmail = to.params.email;
        const routeid = to.params.employeeid;
        if(JSON.parse(window.localStorage.getItem("user")) == null){
            router.push('/unauthorized')
        }
        var id;
        var email = JSON.parse(window.localStorage.getItem("user")).email;
        if(role == 'user' && to.fullPath == '/dashboard'){
            return next()
        }
        EmployeeDataService.retrieveEmployeeByEmail(email).then((res)=>{
            id=res.data.employeeid
            if(
                role == 'admin' ||
                routeEmail === email ||
                routeid === id
            ){
                return next()
            }else{
                router.push('/unauthorized')
            }
        })
    }else{
        return next()
    }



})

export default router;