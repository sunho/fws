import {
  Component,
  OnInit,
  Input,
  ViewEncapsulation,
  Output,
  EventEmitter,
} from '@angular/core';
import { ResListOption } from '../res-list/res-list.component';

@Component({
  selector: 'app-res-card',
  templateUrl: './res-card.component.html',
})
export class ResCardComponent implements OnInit {
  @Output()
  addClick = new EventEmitter();
  @Output()
  detailClick = new EventEmitter();

  @Input()
  options: ResListOption[];
  @Input()
  title: string;
  @Input()
  names: string[];
  @Input()
  keys: string[];
  @Input()
  items: any[];

  constructor() {}

  ngOnInit(): void {}

  onAddClick(): void {
    this.addClick.emit();
  }
}
