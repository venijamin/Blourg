import {Component, OnInit} from '@angular/core';
import {PostService} from "../post-service/post.service";
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
  selector: 'app-post-list',
  templateUrl: './post-list.component.html',
  standalone: true,
  imports: [
    NgForOf
  ],
  styleUrls: ['./post-list.component.css']
})

export class PostListComponent implements OnInit {

  constructor(private postService: PostService) { }
  posts: Post[] = []

  ngOnInit() {
    this.getPosts()
  }

  getPosts() {
    this.postService.getPosts().subscribe(
      value => this.posts = value
    )
  }

  createPost(newPost : Post) {
    this.postService.createPost(newPost).subscribe()
    // refresh the posts
    return this.getPosts()
  }

  deletePost(postId: string) {
    this.postService.deletePost(postId).subscribe()
    // refresh the posts
    return this.getPosts()
  }
}
