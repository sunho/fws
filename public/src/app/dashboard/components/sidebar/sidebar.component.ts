import { SideBarItem } from './../../models/sidebar';
import { Component, OnInit, Input } from '@angular/core';

@Component({
    selector: 'app-sidebar',
    templateUrl: './sidebar.component.html'
})
export class SideBarComponent implements OnInit {
    @Input() current: string;
    items: SideBarItem[] = [
        { title: 'home', icon: 'home', href: '/dashboard'},
        { title: 'build', icon: 'hammer', href: 'build'},
        { title: 'volumes', icon: 'hdd', href: 'build'},
        { title: 'configs', icon: 'document', href: 'build'},
    ];
    constructor() { }

    ngOnInit(): void { }
}
