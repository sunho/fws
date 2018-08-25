import { Component, OnInit } from '@angular/core';

import { AppConfig } from '../../../app.config';

@Component({
  selector: 'app-small-header',
  templateUrl: './small-header.component.html',
  styleUrls: ['./small-header.component.scss']
})
export class SmallHeaderComponent implements OnInit {
  name: string;

  constructor() {
    this.name = AppConfig.siteName;
  }

  ngOnInit() {
  }
}
