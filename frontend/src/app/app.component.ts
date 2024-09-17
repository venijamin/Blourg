import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {PostListComponent} from "./posts/post-list/post-list.component";
import {provideHttpClient, withFetch} from "@angular/common/http";
import {NavBarComponent} from "./nav-bar/nav-bar.component";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, PostListComponent, NavBarComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  title = 'frontend';
}
