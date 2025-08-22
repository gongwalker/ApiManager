/**
 * Markdown编辑器初始化
 * 只用于 name="memo" 的 textarea
 */

let memoEditor = null;

// This function will be called when a file is selected or pasted.
function handleImageUpload(file) {
    if (!memoEditor) {
        console.error("Editor not initialized. Cannot upload image.");
        return;
    }

    layui.use('layer', function() {
        const layer = layui.layer;
        const cm = memoEditor.codemirror;

        if (file.size > 1024 * 1024 * 2) { // 2MB limit
            return layer.msg('Image is too large (max 2MB)', { icon: 2 });
        }

        const formData = new FormData();
        formData.append('image', file);
        const placeholder = `![Uploading ${file.name}...]()`;
        cm.replaceSelection(placeholder);
        const loadingIndex = layer.load(1, { shade: [0.1, '#fff'] });

        fetch('/api/upload', { method: 'POST', body: formData })
            .then(response => {
                if (!response.ok) {
                    return response.json().then(err => { throw new Error(err.msg || 'Upload failed') });
                }
                return response.json();
            })
            .then(result => {
                if (result && result.filename) {
                    const newText = `![${file.name}](${result.filename})`;
                    const doc = cm.getDoc();
                    doc.setValue(doc.getValue().replace(placeholder, newText));
                    layer.msg('Upload successful!', { icon: 1 });
                } else {
                    throw new Error('Invalid server response');
                }
            })
            .catch(error => {
                layer.msg('Upload failed: ' + error.message, { icon: 2 });
                const doc = cm.getDoc();
                doc.setValue(doc.getValue().replace(placeholder, ''));
            })
            .finally(() => {
                layer.close(loadingIndex);
            });
    });
}

function initMarkdownEditors() {
    // 只查找 name="memo" 的 textarea
    const textarea = document.querySelector('textarea[name="memo"]');
    
    // 如果找到了 textarea 且尚未初始化
    if (textarea && textarea.offsetParent !== null && !textarea.simplemde) {
        memoEditor = new SimpleMDE({
            element: textarea,
            spellChecker: false,
            autofocus: false,
            placeholder: textarea.getAttribute('placeholder'),
            toolbar: [
                'bold', 'italic', 'heading', '|',
                'code', 'quote', 'unordered-list', 'ordered-list', '|',
                'link', 
                {
                    name: "image",
                    action: function(editor) {
                        const input = document.createElement('input');
                        input.type = 'file';
                        input.accept = 'image/*';
                        input.style.display = 'none';

                        input.onchange = () => {
                            const file = input.files[0];
                            document.body.removeChild(input);
                            if (file) {
                                handleImageUpload(file);
                            }
                        };
                        document.body.appendChild(input);
                        input.click();
                    },
                    className: "glyphicon glyphicon-picture",
                    title: "Upload Image",
                },
                'table', '|',
                {
                    name: "color",
                    action: function(editor) {
                        const cm = editor.codemirror;
                        const selectedText = cm.getSelection();

                        layui.use('layer', function() {
                            var layer = layui.layer;
                            layer.prompt({
                                formType: 0,
                                value: '#000000',
                                title: 'Enter a color (e.g., #ff0000 or red)',
                                area: ['300px', '150px']
                            }, function(color, index){
                                layer.close(index);
                                if (color) {
                                    const text = '<span style="color:' + color + '">' + selectedText + '</span>';
                                    cm.replaceSelection(text);
                                }
                            });
                        });
                    },
                    className: "glyphicon glyphicon-pencil",
                    title: "Font Color",
                },
                '|',
                'preview', 'side-by-side', 'fullscreen', '|',
                'guide'
            ],
            status: ['lines', 'words'],
        });
        textarea.simplemde = memoEditor; // 标记为已初始化

        // Add the paste handler to the CodeMirror instance
        const cm = memoEditor.codemirror;
        cm.on('paste', function(cmInstance, event) {
            if (event.clipboardData && event.clipboardData.items) {
                for (let i = 0; i < event.clipboardData.items.length; i++) {
                    const item = event.clipboardData.items[i];
                    if (item.type.indexOf('image') !== -1) {
                        event.preventDefault();
                        const blob = item.getAsFile();
                        const file = new File([blob], `screenshot-${Date.now()}.png`, {
                            type: blob.type,
                        });
                        handleImageUpload(file);
                        return; // Stop processing other items
                    }
                }
            }
        });
    }
}

function syncMarkdownEditors() {
    if (memoEditor) {
        const textarea = memoEditor.element;
        textarea.value = memoEditor.value();
    }
    return true;
}

// 挂载到 window 对象
window.initMarkdownEditors = initMarkdownEditors;
window.syncMarkdownEditors = syncMarkdownEditors;