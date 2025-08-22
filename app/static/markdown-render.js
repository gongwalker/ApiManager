/**
 * Markdown渲染器
 * 用于渲染所有带有 .markdown-to-render 类的元素
 */

let renderedMemoElements = new Set();

function renderMarkdownContent() {
    // 查找所有带有 .markdown-to-render 类的 pre 元素
    const elements = document.querySelectorAll('pre.markdown-to-render');

    elements.forEach(element => {
        try {
            // 如果找到了元素且尚未渲染
            if (element && !renderedMemoElements.has(element)) {
                const originalContent = element.textContent || '';
                if (originalContent.trim() === '') {
                    return; // 不渲染空内容
                }

                // 检查 SimpleMDE 是否已加载
                if (typeof SimpleMDE === 'undefined') {
                    console.error('SimpleMDE is not loaded, cannot render markdown.');
                    return;
                }

                // 标记为已处理，防止重复渲染
                renderedMemoElements.add(element);

                // 创建一个临时的SimpleMDE实例来使用其解析器
                const tempElement = document.createElement('textarea');
                // 隐藏临时元素
                tempElement.style.display = 'none';
                document.body.appendChild(tempElement);
                
                const tempEditor = new SimpleMDE({ element: tempElement, toolbar: false });
                
                // 使用实例方法渲染Markdown
                const renderedHTML = tempEditor.markdown(originalContent);
                
                // 卸载并清理临时编辑器和元素
                tempEditor.toTextArea();
                document.body.removeChild(tempElement);
                
                const renderedDiv = document.createElement('div');
                renderedDiv.className = 'markdown-body'; // for styling
                renderedDiv.innerHTML = renderedHTML;
                
                // 替换原始 pre 元素
                if (element.parentNode) {
                    element.parentNode.replaceChild(renderedDiv, element);
                }
            }
        } catch (e) {
            console.error('Error rendering markdown for element:', element, e);
        }
    });
}

// 挂载到 window 对象，以便 common.js 调用
window.renderMarkdownContent = renderMarkdownContent;