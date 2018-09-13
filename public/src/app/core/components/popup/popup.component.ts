import { PopupService } from './../../services/popup.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-popup',
  templateUrl: './popup.component.html',
})
export class PopupComponent implements OnInit {
  msgs: string[] = [];
  constructor(private popupService: PopupService) {}

  ngOnInit(): void {
    this.popupService.$msg.subscribe(
      m => {
        this.msgs.push(m);
      },
      _ => {}
    );
  }

  remove(i: number): boolean {
    this.msgs.splice(i, 1);
    return false;
  }
}
