# How to make an SDK release

Releasing the SDK consists of creating a GitHub release manually,
there is no CI/CD pipeline that automatically builds and publishes the SDK,
there is no publishing to any package manager.

## Prerequisites

- A GitHub account with write access to the repository
- A local clone of the repository

## Steps

1. Make sure you are on the main branch and have the latest changes

   ```bash
   git checkout main
   git pull
   ```

2. Build the SDK

   ```bash
   make
   ```

3. Update the CHANGELOG.md file to document the changes in the new release, following the format shown in the file, and commit the changes

   ```bash
   git add CHANGELOG.md
   git commit -m "Update CHANGELOG for v1.0.16"
   git push
   ```

4. Create a new tag with the appropriate version number (following semantic versioning)

   ```bash
   git tag v1.0.16
   ```

5. Push the tag

   ```bash
   git push origin v1.0.16
   ```

6. Then go to GitHub and create a release based on this tag.
