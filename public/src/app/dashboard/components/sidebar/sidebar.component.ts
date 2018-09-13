import { STRINGS } from '../../../../locale/strings';
import { SideBarItem } from './../../models/sidebar';
import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
})
export class SideBarComponent implements OnInit {
  show = false;
  items: SideBarItem[] = [
    { title: STRINGS.HOME, icon: 'home', href: 'home' },
    { title: STRINGS.VOLUME, icon: 'hdd', href: 'volume' },
    { title: STRINGS.CONFIG, icon: 'document', href: 'config' },
  ];
  constructor() {}

  ngOnInit(): void {}

  onToggleClick(): boolean {
    this.show = !this.show;
    return false;
  }

  onItemClick(): boolean {
    this.show = false;
    // should find another way to make it generic
    document.querySelector('.dash').scrollTo(0, 0);
    return false;
  }
}
