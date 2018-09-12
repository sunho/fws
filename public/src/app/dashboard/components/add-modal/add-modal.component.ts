import { AddModalService, AddModal } from './../../services/add-modal.service';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators, NgForm } from '@angular/forms';

@Component({
    selector: 'app-add-modal',
    templateUrl: './add-modal.component.html'
})
export class AddModalComponent implements OnInit {
    constructor(private addModalService: AddModalService,
        private formBuilder: FormBuilder,
    ) { }

    formGroup: FormGroup;
    current: AddModal;

    ngOnInit(): void {
        this.formGroup = this.formBuilder.group({
            username: ['', Validators.required],
        });
        this.addModalService.$mod.subscribe(
            m => {
                this.current = m;
            },
            _ => { }
        );
    }

    onSubmit(f: NgForm): void {
    }
}
