import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {
  ios = !!navigator.platform && /iPad|iPhone|iPod/.test(navigator.platform);

  ngOnInit(): void {}
}
