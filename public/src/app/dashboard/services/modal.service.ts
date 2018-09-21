import { Injectable } from '@angular/core';
import { Observable, Observer } from 'rxjs';


export interface ModalItem {
  name: string;
  key: string;
  initial?: string;
  disabled?: boolean;
  textfield?: boolean;
}
export interface Modal {
  title: string;
  items: ModalItem[];
  button: string;
  callback: (obj: object) => Observable<boolean>;
}

@Injectable({
  providedIn: 'root',
})
export class ModalService {
  mod: Observer<Modal>;
  $mod: Observable<Modal>;
  constructor() {
    this.$mod = Observable.create(observer => {
      this.mod = observer;
    });
  }

  createMod(mod: Modal): void {
    this.mod.next(mod);
  }
}
