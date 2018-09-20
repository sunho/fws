import { ModalService, Modal } from './../../services/modal.service';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators, NgForm } from '@angular/forms';

@Component({
  selector: 'app-modal',
  templateUrl: './modal.component.html'
})
export class ModalComponent implements OnInit {
  constructor(
    private modalService: ModalService,
    private formBuilder: FormBuilder
  ) {}

  formGroup: FormGroup;
  current: Modal;

  ngOnInit(): void {
    this.modalService.$mod.subscribe(
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
