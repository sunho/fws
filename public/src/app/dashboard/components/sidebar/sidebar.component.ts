import { routeName } from './../../dashboard-routing.module';
import { SideBarItem } from './../../models/sidebar';
import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
})
export class SideBarComponent implements OnInit {
  items: SideBarItem[] = [
    { title: 'home', icon: 'home', href: `./` },
    { title: 'volumes', icon: 'hdd', href: `volumes` },
    { title: 'configs', icon: 'document', href: `configs` },
  ];
  constructor() {}

  ngOnInit(): void {}
}
