import { MonitITPage } from './app.po';

describe('monit-it App', function() {
  let page: MonitITPage;

  beforeEach(() => {
    page = new MonitITPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
