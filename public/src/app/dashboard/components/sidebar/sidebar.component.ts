import { routeName } from "./../../dashboard-routing.module"
import { SideBarItem } from "./../../models/sidebar"
import { Component, OnInit, Input } from "@angular/core"

@Component({
  selector: "app-sidebar",
  templateUrl: "./sidebar.component.html",
})
export class SideBarComponent implements OnInit {
  show: boolean
  items: SideBarItem[] = [
    { title: "home", icon: "home", href: `./` },
    { title: "volume", icon: "hdd", href: `volume` },
    { title: "config", icon: "document", href: `config` },
  ]
  constructor() {}

  ngOnInit(): void {}

  onToggleClick(): void {
    this.show = !this.show
  }

  onItemClick(): void {
    this.show = false
    // should find another way to make it generic
    document.querySelector(".dash").scrollTo(0, 0)
  }
}
