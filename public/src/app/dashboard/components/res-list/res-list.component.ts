import { Component, OnInit, Input } from '@angular/core';

@Component({
    selector: 'app-res-list',
    templateUrl: './res-list.component.html'
})
export class ResListComponent implements OnInit {
    @Input()
    names: string[];
    @Input()
    keys: string[];
    @Input()
    items: any[];
    constructor() { }

    ngOnInit(): void { }
}
