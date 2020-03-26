import { Routes } from '@angular/router'
import { HomeComponent } from './home/home.component'
import { UserHomeComponent } from './user-home/user-home.component'

export const routeArray: Routes = [
    {
        path: '',
        component: HomeComponent
    }, {
        path: '/user/login',
        component: HomeComponent
    }, {
        path: '/user/register',
        component: HomeComponent
    }, {
        path: '/user/:uid',
        component: UserHomeComponent
    }, {
        path: '/user/:uid/category',
        component: UserHomeComponent
    }, {
        path: '/user/:uid/category/:cid',
        component: UserHomeComponent
    }, {
        path: '/user/:uid/category/:cid/bookmark',
        component: UserHomeComponent
    }, {
        path: '/user/:uid/category/:cid/bookmark/:bid',
        component: UserHomeComponent
    }, {
        path: '**',
        component: HomeComponent
    }
]