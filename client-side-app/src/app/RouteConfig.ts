import { Routes } from '@angular/router'
import { HomeComponent } from './components/home/home.component'
import { UserHomeComponent } from './components/user-home/user-home.component'
import { AdminHomeComponent } from './components/admin-home/admin-home.component'

export const routeArray: Routes = [
    {
        path: '',
        component: HomeComponent
    }, {
        path: 'login',
        component: HomeComponent
    }, {
        path: 'register',
        component: HomeComponent
    }, {
        path: ':name/home',
        component: UserHomeComponent
    }, {
        path: ':name/home/category/:categoryID',
        component: UserHomeComponent
    }, {
        path: ':name/admin/home',
        component: AdminHomeComponent
    }, {
        path: '**',
        component: HomeComponent
    }
]