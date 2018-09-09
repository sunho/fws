import { Routes, RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { HomeComponent } from './pages/home/home.component';
import { BuildComponent } from './pages/build/build.component';


const routes: Routes = [
    { path: '', component: HomeComponent },
    { path: 'build', component: BuildComponent},
];


@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class DashBoardRoutingModule {}
