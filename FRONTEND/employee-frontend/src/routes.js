import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

const router = new Router({
    mode: 'history',
    routes: [
        {
            path: "/employees",
            name: "Employees",
            component: () => import("./components/Employees"),
        },
        {
            path: "/employee/new",
            name: "NewEmployee",
            component: () => import("./components/NewEmployee"),
        },
        {
            path: "/employee/:emailid",
            name: "Employee",
            component: () => import("./components/Employee"),
        },
        {
            path: "/trainings",
            name: "Trainings",
            component: () => import("./components/Trainings"),
        },
        {
            path: "/training/new",
            name: "NewTraining",
            component: () => import("./components/NewTraining"),
        },
        {
            path: "/training/:courseCode",
            name: "Training",
            component: () => import("./components/Training"),
        },      
    ]
});

export default router;