package nestor

type parser struct {
}

// Parse starts and looks for tutorial packages it can find.
// A Tutorial must have the following layout:
//  - some dir
//  + some
//    + other dir
//      + 01-tutorial-0 (Tutorial.ID)
//        - tutorial.yml (contains Tutorial.Title and Tutorial.Subtitle)
//        - image.png/jpg (Tutorial.TeaserUrl = Tutorial.ID + image.png)
//        + 01-chapter-0
//        + 02-chapter-1 (Chapter.ID)
//          - image.png (Chapter.TeaserUrl = Tutorial.ID + Chapter.ID + image.png)
//          - chapter.md (
func Parse(rootDir string) ([]Tutorial, error) {

}
