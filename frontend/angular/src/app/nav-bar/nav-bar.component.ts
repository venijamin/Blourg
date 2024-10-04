import { Component } from '@angular/core';
import {PostListComponent} from "../posts/post-list/post-list.component";
import {RouterLink, RouterLinkActive} from "@angular/router";
import {PostFormComponent} from "../posts/post-form/post-form.component";

@Component({
  selector: 'app-nav-bar',
  standalone: true,
  imports: [
    PostListComponent,
    RouterLink,
    RouterLinkActive,
    PostFormComponent
  ],
  templateUrl: './nav-bar.component.html',
  styleUrl: './nav-bar.component.css'
})
export class NavBarComponent {
  sidebarActive: boolean = false;

  toggleSidebar() {
    this.sidebarActive = !this.sidebarActive;
  }
}
