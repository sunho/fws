import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { DropdownItem } from '../../../core/components/form-dropdown/form-dropdown.component';

export interface ResListOption {
  title: string;
  func: (any, string) => void;
}

@Component({
  selector: 'app-res-list',
  templateUrl: './res-list.component.html',
})
export class ResListComponent implements OnInit {
  @Output()
  itemClick = new EventEmitter();

  @Input()
  options: ResListOption[];
  @Input()
  names: string[];
  @Input()
  keys: string[];
  @Input()
  items: any[];
  constructor() {}

  ngOnInit(): void {}

  createDrops(obj: any): DropdownItem[] {
    return this.options.map(o => {
      return {
        title: o.title,
        func: (str: string) => {
          o.func(obj, str);
        },
      };
    });
  }
}
