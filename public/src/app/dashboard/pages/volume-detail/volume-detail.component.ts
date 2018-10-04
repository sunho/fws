import { Volume } from './../../models/bot';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
    selector: 'app-volume-detail',
    templateUrl: './volume-detail.component.html'
})
export class VolumeDetailComponent implements OnInit {
    current: Volume;

    constructor(
        private route: ActivatedRoute,
    ) { }

    ngOnInit(): void {
        this.route.data.subscribe(
            data => {
                this.current = data['volume'];
            },
            _ => { }
        );
    }
}
