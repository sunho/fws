import { Injectable } from '@angular/core';
import { Observable, Observer } from 'rxjs';

export interface AddModal {
    names: string[];
    keys: string[];
    callback: (obj: object) => Observable<boolean>;
}

@Injectable({
  providedIn: 'root',
})
export class AddModalService {
  mod: Observer<AddModal>;
  $mod: Observable<AddModal>;
  constructor() {
    this.$mod = Observable.create(observer => {
      this.mod = observer;
    });
  }

  createMod(mod: AddModal): void {
    this.mod.next(mod);
  }
}
