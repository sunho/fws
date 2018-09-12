import { AddModalComponent } from './components/add-modal/add-modal.component';
import { ResCardComponent } from './components/res-card/res-card.component';
import { ResListComponent } from './components/res-list/res-list.component';
import { ConfigComponent } from './pages/config/config.component';
import { BotSelectComponent } from './components/bot-select/bot-select.component';
import { HeaderComponent } from './components/header/header.component';
import { RouterModule } from '@angular/router';
import { SideBarComponent } from './components/sidebar/sidebar.component';
import { DashBoardRoutingModule } from './dashboard-routing.module';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './pages/home/home.component';
import { CoreModule } from '../core/core.module';
import { DashComponent } from './pages/dash/dash.component';
import { VolumeComponent } from './pages/volume/volume.component';
import { UserProfileComponent } from './components/user-profile/user-profile.component';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    AddModalComponent,
    ResListComponent,
    ResCardComponent,
    UserProfileComponent,
    VolumeComponent,
    ConfigComponent,
    BotSelectComponent,
    DashComponent,
    HeaderComponent,
    HomeComponent,
    SideBarComponent,
  ],
  imports: [DashBoardRoutingModule, ReactiveFormsModule, FormsModule, RouterModule, CoreModule, CommonModule],
  exports: [],
  providers: [],
})
export class DashBoardModule {}
