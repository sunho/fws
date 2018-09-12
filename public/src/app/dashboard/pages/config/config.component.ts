import { Component, OnInit } from '@angular/core';
import { Config, Env, Bot } from '../../models/bot';

@Component({
  selector: 'app-config',
  templateUrl: './config.component.html',
})
export class ConfigComponent implements OnInit {
  current: Bot;

  envNames = [
    'name',
    'value',
  ];

  envKeys = [
    'name',
    'path',
  ];

  envItems: Env[];

  confNames = [
    'name',
    'path',
  ];

  confKeys = [
    'name',
    'path',
  ];

  confItems: Config[];

  constructor() {}

  ngOnInit(): void {}
}
