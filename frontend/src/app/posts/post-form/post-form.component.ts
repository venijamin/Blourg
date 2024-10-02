import {Component, OnInit} from '@angular/core';
import {PostService} from "../post-service/post.service";
import {Post} from "../../model/post";
import {MarkdownModule} from "ngx-markdown";
import {FormsModule} from "@angular/forms";

@Component({
  selector: 'app-post-form',
  standalone: true,
  imports: [MarkdownModule, FormsModule],
  templateUrl: './post-form.component.html',
  styleUrl: './post-form.component.css'
})


export class PostFormComponent implements OnInit {
  username: string = '';
  title: string = '';
  body: string = '';

  constructor(private postService: PostService) { }

  createPost() {
    const newPost = {
      username: this.username,
      title: this.title,
      body: this.body,
    };

    this.postService.createPost(newPost).subscribe()
  }

  deletePost(postId: string) {
    this.postService.deletePost(postId).subscribe()
  }

  ngOnInit(): void {
  }

}
