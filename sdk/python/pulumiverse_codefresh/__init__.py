# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .account import *
from .context import *
from .pipeline import *
from .project import *
from .provider import *
from .registry import *
from .team import *
from .user import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumiverse_codefresh.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumiverse_codefresh.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "codefresh",
  "mod": "index/account",
  "fqn": "pulumiverse_codefresh",
  "classes": {
   "codefresh:index/account:Account": "Account"
  }
 },
 {
  "pkg": "codefresh",
  "mod": "index/context",
  "fqn": "pulumiverse_codefresh",
  "classes": {
   "codefresh:index/context:Context": "Context"
  }
 },
 {
  "pkg": "codefresh",
  "mod": "index/pipeline",
  "fqn": "pulumiverse_codefresh",
  "classes": {
   "codefresh:index/pipeline:Pipeline": "Pipeline"
  }
 },
 {
  "pkg": "codefresh",
  "mod": "index/project",
  "fqn": "pulumiverse_codefresh",
  "classes": {
   "codefresh:index/project:Project": "Project"
  }
 },
 {
  "pkg": "codefresh",
  "mod": "index/registry",
  "fqn": "pulumiverse_codefresh",
  "classes": {
   "codefresh:index/registry:Registry": "Registry"
  }
 },
 {
  "pkg": "codefresh",
  "mod": "index/team",
  "fqn": "pulumiverse_codefresh",
  "classes": {
   "codefresh:index/team:Team": "Team"
  }
 },
 {
  "pkg": "codefresh",
  "mod": "index/user",
  "fqn": "pulumiverse_codefresh",
  "classes": {
   "codefresh:index/user:User": "User"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "codefresh",
  "token": "pulumi:providers:codefresh",
  "fqn": "pulumiverse_codefresh",
  "class": "Provider"
 }
]
"""
)
