import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-icon',
  templateUrl: './icon.component.html',
  styleUrls: ['./icon.component.scss'],
})
export class IconComponent implements OnInit {
  @Input()
  icon: string;

  base = document.location.origin;

  constructor() {}

  ngOnInit(): void {}
}
