import {Component, OnInit} from '@angular/core';
import {PostService} from "../post.service";
import {NgForOf} from "@angular/common";

interface Post {
  post_id: string;
  username: string;
  title: string;
  body: string;
  up_vote: number;
  down_vote: number;
}

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  standalone: true,
  imports: [
    NgForOf
  ],
  styleUrls: ['./post.component.css']
})

export class PostComponent implements OnInit {
  posts: Post[] = [];

  newPost: Post = {
post_id: '',
    body: '',
    down_vote: 0,
    up_vote: 2,
    username: '',
    title: '',

  };

  constructor(private postService: PostService) { }

  ngOnInit() {
    this.getPosts();
  }

  getPosts() {
    this.postService.getPosts()
      .subscribe(
        (response) => {
          this.posts = response;
        },
        (error) => {
          console.error(error);
        }
      );
  }

  createPost() {
    this.postService.createPost(this.newPost)
      .subscribe(
        (response) => {
          console.log('Post created:', response);
          // Refresh the user list after creating a new user
          this.getPosts();
        },
        (error) => {
          console.error(error);
        }
      );
  }

  deletePost(postId: string) {
    this.postService.deletePost(postId)
      .subscribe(
        () => {
          console.log('Post deleted');
          // Refresh the user list after deleting a user
          this.getPosts();
        },
        (error) => {
          console.error(error);
        }
      );
  }
}
