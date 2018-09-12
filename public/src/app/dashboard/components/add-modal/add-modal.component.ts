import { AddModalService, AddModal } from './../../services/add-modal.service';
import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'app-add-modal',
    templateUrl: './add-modal.component.html'
})
export class AddModalComponent implements OnInit {
    constructor(private addModalService: AddModalService) { }

    current: AddModal;

    
    ngOnInit(): void {
        this.addModalService.$mod.subscribe(
            m => {
                this.current = m;
            },
            _ => { }
        );
    }
}
