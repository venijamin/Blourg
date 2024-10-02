import {Component, OnInit} from '@angular/core';
import {PostService} from "../post-service/post.service";
import {NgForOf} from "@angular/common";
import {Post} from "../../model/post";



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


}
