import axios from 'axios';

const BaseURL = 'https://api.micro.mu/distributed'

export default async function Call(path: string, params?: any): Promise<any> {
  return axios.get(BaseURL + path, params)
}

export class Note {
  id: string;
  title: string;
  text: string;
  created: Date;

  constructor(args: any) {
    this.id = args.id;
    this.title = args.title;
    this.text = args.text;
    this.created = new Date(parseInt(args.created) * 1000);
  }
}