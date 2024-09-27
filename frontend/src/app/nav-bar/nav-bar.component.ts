import { Component } from '@angular/core';
import {PostListComponent} from "../posts/post-list/post-list.component";

@Component({
  selector: 'app-nav-bar',
  standalone: true,
  imports: [
    PostListComponent
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
