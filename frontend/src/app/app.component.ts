import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {PostComponent} from "./post/post.component";
import {provideHttpClient, withFetch} from "@angular/common/http";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, PostComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  title = 'frontend';
}
