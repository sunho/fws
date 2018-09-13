import { AddModalService, AddModal } from './../../services/add-modal.service';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators, NgForm } from '@angular/forms';

@Component({
  selector: 'app-add-modal',
  templateUrl: './add-modal.component.html',
})
export class AddModalComponent implements OnInit {
  constructor(
    private addModalService: AddModalService,
    private formBuilder: FormBuilder
  ) {}

  formGroup: FormGroup;
  current: AddModal;

  ngOnInit(): void {
    this.addModalService.$mod.subscribe(
      m => {
        this.current = m;
        const obj = {};
        for (const key of this.current.keys) {
          obj[key] = ['', Validators.required];
        }
        this.formGroup = this.formBuilder.group(obj);
      },
      _ => {}
    );
  }

  onCancelClick(): boolean {
    this.current = null;
    return false;
  }

  onSubmit(f: NgForm): void {
    if (f.valid) {
      this.current.callback(f.value).subscribe(
        b => {
          if (b) {
            this.current = null;
          }
        },
        _ => {}
      );
    }
  }
}
