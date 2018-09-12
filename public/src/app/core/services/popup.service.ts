import { Injectable } from '@angular/core';
import { Observable, Observer } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class PopupService {
  msg: Observer<string>;
  $msg: Observable<string>;
  constructor() {
    this.$msg = Observable.create(observer => {
      this.msg = observer;
    });
  }

  createMsg(str: string): void {
    this.msg.next(str);
  }
}
