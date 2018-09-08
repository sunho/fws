import { DashBoardRoutingModule } from './dashboard-routing.module';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './pages/home/home.component';

@NgModule({
    declarations: [
        HomeComponent
    ],
    imports: [
        DashBoardRoutingModule,
        CommonModule
    ],
    exports: [],
    providers: [],
})
export class DashBoardModule { }
