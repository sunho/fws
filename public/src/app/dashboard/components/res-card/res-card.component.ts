import { Component, OnInit, Input, ViewEncapsulation } from '@angular/core';

@Component({
    selector: 'app-res-card',
    templateUrl: './res-card.component.html'
})
export class ResCardComponent implements OnInit {
    @Input()
    addroute: string;
    @Input()
    detailroute: string;

    @Input()
    title: string;
    @Input()
    names: string[];
    @Input()
    keys: string[];
    @Input()
    items: any[];

    constructor() { }

    ngOnInit(): void { }
}
