import { AuthGuardService } from './../core/services/auth-guard.service';
import { Routes, RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { BuildComponent } from './pages/build/build.component';
import { DashComponent } from './pages/dash/dash.component';
import { HomeComponent } from './pages/home/home.component';

const routes: Routes = [
    {
        path: '',
        component: DashComponent,
        children: [
            { path: '', component: HomeComponent },
            { path: 'build', component: BuildComponent },
        ]
    },
];

export const routeName = 'dashboard';

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class DashBoardRoutingModule {}
