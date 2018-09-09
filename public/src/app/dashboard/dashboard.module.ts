import { HeaderComponent } from './components/header/header.component';
import { RouterModule } from '@angular/router';
import { SideBarComponent } from './components/sidebar/sidebar.component';
import { DashBoardRoutingModule } from './dashboard-routing.module';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './pages/home/home.component';
import { CoreModule } from '../core/core.module';
import { BuildComponent } from './pages/build/build.component';

@NgModule({
    declarations: [
        HeaderComponent,
        BuildComponent,
        HomeComponent,
        SideBarComponent
    ],
    imports: [
        DashBoardRoutingModule,
        RouterModule,
        CoreModule,
        CommonModule
    ],
    exports: [],
    providers: [],
})
export class DashBoardModule { }
