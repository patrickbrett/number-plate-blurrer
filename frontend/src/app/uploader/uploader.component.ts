import { Component, OnInit } from '@angular/core';
import axios from 'axios';

@Component({
  selector: 'app-uploader',
  templateUrl: './uploader.component.html',
  styleUrls: ['./uploader.component.scss'],
})
export class UploaderComponent implements OnInit {
  upload: any = {};
  filename = '';
  uploadUrl = '';
  apiBase = 'https://pn53za3axf.execute-api.ap-southeast-2.amazonaws.com/Prod';

  constructor() {}

  ngOnInit(): void {}

  async getUploadUrl() {
    const response = (await axios.put(`${this.apiBase}/images`)).data;
    console.log(response);

    this.uploadUrl = response.uploadUrl;
    this.filename = response.filename;
  }

  setFile = (e) => {
    const { target } = e;
    const file = target?.files && target.files[0];

    if (!file) return;

    var url = URL.createObjectURL(file);
    var img = new Image();
    img.src = url;

    img.onload = () => {
      this.upload.src = img.src;
      this.upload.file = file;
    };
  };

  submitImage = async (e) => {
    e.preventDefault();

    const { file } = this.upload;

    try {
      await axios.put(this.uploadUrl, file, {
        headers: { 'Content-Type': 'image/jpeg' },
      });

      console.log('uploaded!');
    } catch (err) {
      console.error(err, err.message);
    }
  };

  loadLabels = async () => {
    const response = (await axios.get(`${this.apiBase}/labels?filename=${this.filename}`)).data;

    console.log(response);
  };
}
