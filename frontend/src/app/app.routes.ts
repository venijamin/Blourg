import { Routes } from '@angular/router';
import {AppComponent} from "./app.component";
import {PostFormComponent} from "./posts/post-form/post-form.component";

export const routes: Routes = [
  { path: '/posts/create', component: PostFormComponent }, // Route for creating a post

];
